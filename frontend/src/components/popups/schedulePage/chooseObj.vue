<template>
    <div class="popup popup-squared">
        <div class="popup__header">
            <h2 class="h2">Выбрать объект</h2>
            <button class="btn-no-bg" id="popupCloseBtn">
                <svg width="51" height="51" viewBox="0 0 51 51" fill="none" xmlns="http://www.w3.org/2000/svg">
                    <path d="M12.75 12.75L38.25 38.25" stroke="#A6A6A6" stroke-width="3" stroke-linecap="round" />
                    <path d="M38.25 12.75L12.75 38.25" stroke="#A6A6A6" stroke-width="3" stroke-linecap="round" />
                </svg>
            </button>
        </div>
        <div class="popup__main">
            <form action="" method="" class="popup__form">
                <div class="popup__form_element">
                    <p class="text">Тип объекта</p>
                    <div class="select required" id="scheduleType" required>
                        <button class="btn-no-bg select__btn" type="button" name="scheduleType" value="" data-index="-1"
                            data-select="toggle">
                            Выбрать тип объекта
                            <svg width="15" height="16" viewBox="0 0 15 16" fill="none"
                                xmlns="http://www.w3.org/2000/svg">
                                <path d="M0.75 2.3999L7.5 9.5999L14.25 2.3999" stroke="#A6A6A6" stroke-width="1.5"
                                    stroke-linecap="round" stroke-linejoin="round" />
                            </svg>
                        </button>
                        <div class="select__dropdown">
                            <ul class="select__options" data-index="0">
                                <li class="select__option" data-index="1" data-value="groups" data-select="option">
                                    Группы</li>
                                <li class="select__option" data-index="2" data-value="subjects" data-select="option">
                                    Предметы</li>
                                <li class="select__option" data-index="2" data-value="cabinets" data-select="option">
                                    Кабинеты</li>
                                <li class="select__option" data-index="2" data-value="teachers" data-select="option">
                                    Преподаватели</li>
                            </ul>
                        </div>
                    </div>
                </div>
                <div class="popup__form_element">
                    <p class="text">Неделя</p>
                    <div class="select select-disabled" id="scheduleWeek">
                        <button class="btn-no-bg select__btn" type="button" name="scheduleWeek" value="" data-index="-1"
                            data-select="toggle">
                            Выбрать неделю
                            <svg width="15" height="16" viewBox="0 0 15 16" fill="none"
                                xmlns="http://www.w3.org/2000/svg">
                                <path d="M0.75 2.3999L7.5 9.5999L14.25 2.3999" stroke="#A6A6A6" stroke-width="1.5"
                                    stroke-linecap="round" stroke-linejoin="round" />
                            </svg>
                        </button>
                        <div class="select__dropdown">
                            <ul class="select__options" data-index="0">
                                <li class="select__option" data-index="1" data-value="week1" data-select="option">Чётная
                                </li>
                                <li class="select__option" data-index="2" data-value="week2" data-select="option">
                                    Нечётная</li>
                            </ul>
                        </div>
                    </div>
                </div>
                <div class="popup__form_element">
                    <p class="text">Конкретный объект</p>
                    <div class="select select-disabled" id="scheduleCurrentObj">
                        <button class="btn-no-bg select__btn" type="button" name="scheduleCurrentObj" value=""
                            data-index="-1" data-select="toggle">
                            Выбрать объект
                            <svg width="15" height="16" viewBox="0 0 15 16" fill="none"
                                xmlns="http://www.w3.org/2000/svg">
                                <path d="M0.75 2.3999L7.5 9.5999L14.25 2.3999" stroke="#A6A6A6" stroke-width="1.5"
                                    stroke-linecap="round" stroke-linejoin="round" />
                            </svg>
                        </button>
                        <div class="select__dropdown">
                            <ul class="select__options" data-index="0">

                            </ul>
                        </div>
                    </div>
                </div>

                <div class="popup__footer">
                    <button class="btn btn-disabled" type="submit">
                        Применить
                    </button>
                </div>
            </form>
        </div>

    </div>
</template>