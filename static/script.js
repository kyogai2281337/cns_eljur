function getCookie(name) {
  const cookies = document.cookie.split(';');
  for (let i = 0; i < cookies.length; i++) {
      const cookie = cookies[i].trim();
      // Проверяем, начинается ли кука с искомого имени
      if (cookie.startsWith(name + '=')) {
          // Возвращаем значение куки
          return cookie.substring(name.length + 1);
      }
  }
  // Если кука не найдена, возвращаем пустую строку или можно вернуть null
  return '';
}

// fetch == request in js; Тестовые трекеры на базюклу и прочек сейва куков через фетч
const userblock = document.querySelector(".user")
async function rfetch(url, dataset) {
    try {
        const res = await fetch(url, {
          method: "post",
          headers: {
                    'Accept': 'application/json',
                    'Content-Type': 'text/plain; charset=utf-8',
                    //'Origin': 'google.com'
                  },
                  //make sure to serialize your JSON body
          body: dataset
        });
        //console.log(res)
        const users = await res.json();
        console.log(users);
        //console.log(Array.from(res.headers.entries()));
        console.log(getCookie('journal_auth'));
        }
    catch (err) {
        console.log("Error: ", err);
    }
}

const data = {
  email: "user@example.org",
  password: "Erunda2291337",
  first: "Negr",
  last: "Pidorov"
};

rfetch('http://localhost:6987/signup', JSON.stringify(data));

