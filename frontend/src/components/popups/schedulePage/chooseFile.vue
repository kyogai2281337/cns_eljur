<template>
    <div class="popup-horisontal popup">
        <div class="popup__header">
            <h2 class="h2">
                Выбрать файл
            </h2>
            <button class="btn-no-bg" id="popupCloseBtn">
                <svg width="51" height="51" viewBox="0 0 51 51" fill="none" xmlns="http://www.w3.org/2000/svg">
                    <path d="M12.75 12.75L38.25 38.25" stroke="#A6A6A6" stroke-width="3" stroke-linecap="round" />
                    <path d="M38.25 12.75L12.75 38.25" stroke="#A6A6A6" stroke-width="3" stroke-linecap="round" />
                </svg>
            </button>
        </div>

        <div class="popup__main popup__choose_file">

            <div class="popup__main__wrapper">

                <div class="choose_file__select">
                    <div class="select">
                        <button class="btn-no-bg select__btn" type="button" name="scheduleName" value="" data-index="-1"
                            data-select="toggle">
                            Выбрать расписание
                            <svg width="15" height="16" viewBox="0 0 15 16" fill="none"
                                xmlns="http://www.w3.org/2000/svg">
                                <path d="M0.75 2.3999L7.5 9.5999L14.25 2.3999" stroke="#A6A6A6" stroke-width="1.5"
                                    stroke-linecap="round" stroke-linejoin="round" />
                            </svg>
                        </button>

                        <div class="dropdown__wrapper dropdown-relative">
                            <div class="select__dropdown">
                                <ul class="select__options" data-index="0">
                                    <li class="select__option" data-index="1" data-value="Schedule_1"
                                        data-select="option">Schedule_1</li>
                                    <li class="select__option" data-index="2" data-value="Schedule_2"
                                        data-select="option">Schedule_2</li>
                                    <li class="select__option" data-index="2" data-value="Schedule_2"
                                        data-select="option">Schedule_2</li>
                                    <li class="select__option" data-index="2" data-value="Schedule_2"
                                        data-select="option">Schedule_2</li>
                                </ul>
                            </div>
                        </div>
                    </div>
                </div>
                <div class="choose_file__upload btn btn-bordered btn-bordered__blue">
                    Загрузить с устройства
                    <input type="file" name="schedule" accept=".xlsx, .xls" />
                    <svg width="35" height="35" viewBox="0 0 35 35" fill="none" xmlns="http://www.w3.org/2000/svg">
                        <path
                            d="M16.0417 16.0418C16.0417 15.2364 16.6946 14.5835 17.5001 14.5835C18.3055 14.5835 18.9584 15.2364 18.9584 16.0418V18.9585H21.8751C22.6805 18.9585 23.3334 19.6114 23.3334 20.4168C23.3334 21.2223 22.6805 21.8752 21.8751 21.8752H18.9584V24.7918C18.9584 25.5973 18.3055 26.2502 17.5001 26.2502C16.6946 26.2502 16.0417 25.5973 16.0417 24.7918V21.8752H13.1251C12.3197 21.8752 11.6667 21.2223 11.6667 20.4168C11.6667 19.6114 12.3197 18.9585 13.1251 18.9585H16.0417V16.0418Z"
                            fill="white" />
                        <path fill-rule="evenodd" clip-rule="evenodd"
                            d="M1.45825 5.8335C1.45825 3.41726 3.41701 1.4585 5.83325 1.4585H11.3223C12.9794 1.4585 14.4944 2.39476 15.2354 3.87694L16.5398 6.48568C16.7868 6.97975 17.2918 7.29183 17.8442 7.29183H29.1666C31.5829 7.29183 33.5416 9.25059 33.5416 11.6668V29.1668C33.5416 31.5831 31.5829 33.5418 29.1666 33.5418H5.83325C3.41701 33.5418 1.45825 31.5831 1.45825 29.1668V5.8335ZM5.83325 4.37516C5.02784 4.37516 4.37492 5.02809 4.37492 5.8335V29.1668C4.37492 29.9723 5.02784 30.6252 5.83325 30.6252H29.1666C29.972 30.6252 30.6249 29.9723 30.6249 29.1668V11.6668C30.6249 10.8614 29.972 10.2085 29.1666 10.2085H17.8442C16.1871 10.2085 14.6722 9.27223 13.9311 7.79005L12.6267 5.18131C12.3797 4.68725 11.8747 4.37516 11.3223 4.37516H5.83325Z"
                            fill="white" />
                    </svg>
                </div>

            </div>
            
        </div>

        <div class="popup__footer">
            <button class="btn btn-disabled">
                Применить
            </button>
        </div>

    </div>
</template>