<template>
     <div class="popup popup-horisontal" id="createSchedulePopup">
        <div class="popup__header">
            <p class="h2 text text-warning">
                Создать расписание
            </p>
            <button class="btn-no-bg" id="popupCloseBtn">
                <svg width="51" height="51" viewBox="0 0 51 51" fill="none" xmlns="http://www.w3.org/2000/svg">
                  <path d="M12.75 12.75L38.25 38.25" stroke="#A6A6A6" stroke-width="3" stroke-linecap="round"/>
                  <path d="M38.25 12.75L12.75 38.25" stroke="#A6A6A6" stroke-width="3" stroke-linecap="round"/>
                </svg>              
            </button>
        </div>

        <div class="popup__main">
                <form action="" class="form" id="createScheduleForm">
                    <div class="form__element">
                        <label for="getScheduleName" class="form__label">Название нового расписания:</label>
                        <input type="text" class="form__inp" id="getScheduleName" placeholder="Schedule_1" required> 
                    </div>
                    <div class="form__element">
                        <label for="weeks" class="form__label">Количество недель:</label>
                        <input type="number" class="form__inp" id="weeks" min="1" max="4" placeholder="0" required>
                    </div>
                    <div class="form__element">
                        <label for="maxDays" class="form__label">Общеее количество дней в расписании:</label>
                        <input type="number" class="form__inp" id="maxDays" min="1" max="21" placeholder="0" required>
                    </div>
                    <div class="form__element">
                        <label for="maxDaysPerWeek" class="form__label">Максимальнное количество дней в неделе:</label>
                        <input type="number" class="form__inp" id="maxDaysPerWeek" min="1" max="6" placeholder="0" required>
                    </div>
                    <div class="form__element">
                        <label for="maxPairsPerDay" class="form__label">Количество недель:</label>
                        <input type="number" class="form__inp" id="maxPairsPerDay" min="1" max="7" placeholder="0" required>
                    </div>

                <button id="commit" class="btn btn-bordered button btn btn-bordered btn-bordered__green" method="post">
                    Сгенерировать
                    <svg width="22" height="22" viewBox="0 0 22 22" fill="none" xmlns="http://www.w3.org/2000/svg">
                        <line x1="6.01037" y1="5.65674" x2="15.5563" y2="15.2027" stroke="#5ab17d" stroke-width="1.5" stroke-linecap="round"/>
                        <line x1="5.65674" y1="15.2025" x2="15.2027" y2="5.65658" stroke="#5ab17d" stroke-width="1.5" stroke-linecap="round"/>
                    </svg>        
                </button>
            </form>
        </div>
    </div>
</template>
