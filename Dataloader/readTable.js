const xlsx = require("xlsx");
const path = require("path");

const subjects = require('./subjects.json');

function parseExcel(filePath) {
    const workbook = xlsx.readFile(filePath);
    const sheetName = workbook.SheetNames[0];
    const sheet = workbook.Sheets[sheetName];

    const data = xlsx.utils.sheet_to_json(sheet, { header: 1 });
    const result = [];

    for (let i = 2; i < data.length; i++) {
        const row = data[i];
        if (!row[0] || !row[1]) continue;

        const entry = {
            Преподаватель: row[0],
            Группа: row[1],
            Предмет: row[3],
            "1 полугодие": {
                "Учебных недель": row[17] ? row[4] : "",
                "Пар в неделю": row[17] ? row[5] ? (Math.floor(parseFloat(row[5])) || 1) : "" : "",
                "В нагрузку": row[17] || "",
            },
            "2 полугодие": {
                "Учебных недель": row[31] ? row[18] || "" : "",
                "Пар в неделю": row[31] ? row[19] ? (Math.floor(parseFloat(row[19])) || 1) : "" : "",
                "В нагрузку": row[31] || "",
            },
        };
        result.push(entry);
    }

    return result;
}

const filePath = path.resolve(__dirname, "Сип.xlsx");
const data = parseExcel(filePath); //читается файл с задаными ключами для форматирования json

function groupSubjects(subjects, data) {
    const groups = {};

    data.forEach(entry => {
        const group = entry["Группа"];
        if (!groups[group]) {
            const year = parseInt(group.split('-')[group.split('-').length-1])-1;
            const course = 24 - year;
            groups[group] = {
                id: Object.keys(groups).length + 1,
                name: group,
                course: course,
                short_plan_1year: {},
                short_plan_2year: {}
            };
        }

        const subjectId = subjects.find(s => s.name === entry["Предмет"])?.id;
        if (!subjectId) return;

        const semester1 = entry["1 полугодие"]["Пар в неделю"];
        const semester2 = entry["2 полугодие"]["Пар в неделю"];

        if (semester1) {
            groups[group].short_plan_1year[subjectId] = semester1;
        }
        if (semester2) {
            groups[group].short_plan_2year[subjectId] = semester2;
        }
    });

    return Object.values(groups);
}

const groups = groupSubjects(subjects, data); //объединяет предметы группы в однин объект

function splitGroupData(groups) {
    let idCounter = 1;
    return groups.flatMap(group => {
        const result = [];
        if (Object.keys(group.short_plan_1year).length > 0) {
          result.push({
            id: idCounter++,
            name: `${group.name}-1семестр`,
            course: group.course,
            short_plan: group.short_plan_1year
          });
        }
        if (Object.keys(group.short_plan_2year).length > 0) {
          result.push({
            id: idCounter++,
            name: `${group.name}-2семестр`,
            course: group.course,
            short_plan: group.short_plan_2year
          });
        }
        return result;
    });
}

const output = splitGroupData(groups); //разделяем объект группы на 2 объекта с разными семестрами для специализаций по семестрам

require('fs').writeFileSync('specializations.json', JSON.stringify(output, null, 2));

function generateGroupsLinkSpec(groups) {
    let specializationIdCounter = 1;
    const specializationData = [];
  
    groups.forEach(item => {
        if (Object.keys(item.short_plan).length > 0) {
            specializationData.push({
              id: specializationIdCounter++,
              specialization: {
                id: specializationIdCounter - 1
              },
              name: item.name,
              max_pairs: 18
            });
        }
    });
  
    return specializationData;
}

const output2 = generateGroupsLinkSpec(output); //создаем массив групп с привязкой к специализациям

require('fs').writeFileSync('groups.json', JSON.stringify(output2, null, 2));

function generateTeachersLinkgroupsAndSpecs(groupsSubjects, SubjectsDataFromExcel) {
    let teachers = [];

    groupsSubjects = groupsSubjects.map(schObj => {
        return {
            name: schObj.name,
            id: schObj.id,
        };
    });

    SubjectsDataFromExcel = SubjectsDataFromExcel.map(groupObj => {
        return {
            teacher: groupObj.Преподаватель,
            group: groupObj.Группа,
            subject: groupObj.Предмет,
            subjectId: subjects.find(s => s.name === groupObj.Предмет)?.id,
            year1: groupObj["1 полугодие"]["Пар в неделю"] ? true : false,
            year2: groupObj["2 полугодие"]["Пар в неделю"] ? true : false,
        };
    });

    groupsSubjects.forEach(group => {
        const groupName = group.name.replace("-1семестр", "").replace("-2семестр", "");
        const groupId = group.id;

        const teachersSubjects = SubjectsDataFromExcel.filter(obj => obj.group === groupName);

        teachersSubjects.forEach(subject => {
            let teacherIndex = teachers.findIndex(teacher => teacher.name === subject.teacher);

            if (teacherIndex === -1) {
                teachers.push({
                    name: subject.teacher,
                    capacity: 18,
                    links: {},
                });
                teacherIndex = teachers.length - 1;
            }

            if (subject.year1 && group.name.includes("-1семестр")) {
                const semesterGroupId = `${groupId}`;
                if (!teachers[teacherIndex].links[semesterGroupId]) {
                    teachers[teacherIndex].links[semesterGroupId] = [];
                }
                if (!teachers[teacherIndex].links[semesterGroupId].includes(subject.subjectId)) {
                    teachers[teacherIndex].links[semesterGroupId].push(subject.subjectId);
                }
            }

            if (subject.year2 && group.name.includes("-2семестр")) {
                const semesterGroupId = `${groupId}`;
                if (!teachers[teacherIndex].links[semesterGroupId]) {
                    teachers[teacherIndex].links[semesterGroupId] = [];
                }
                if (!teachers[teacherIndex].links[semesterGroupId].includes(subject.subjectId)) {
                    teachers[teacherIndex].links[semesterGroupId].push(subject.subjectId);
                }
            }
        });
    });

    return teachers;
}

require('fs').writeFileSync('teachers.json', JSON.stringify(generateTeachersLinkgroupsAndSpecs(output2, data), null, 2));