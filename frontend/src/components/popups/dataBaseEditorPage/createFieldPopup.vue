<template>
  <div class="popup popup-horisontal">
    <div class="popup__header">
      <p class="h2 text text-warning">Создать новую запись</p>
      <button class="btn-no-bg" id="popupCloseBtn">
        <svg
          width="51"
          height="51"
          viewBox="0 0 51 51"
          fill="none"
          xmlns="http://www.w3.org/2000/svg"
        >
          <path
            d="M12.75 12.75L38.25 38.25"
            stroke="#A6A6A6"
            stroke-width="3"
            stroke-linecap="round"
          />
          <path
            d="M38.25 12.75L12.75 38.25"
            stroke="#A6A6A6"
            stroke-width="3"
            stroke-linecap="round"
          />
        </svg>
      </button>
    </div>
    <div class="popup__main">
      <form action="" class="form" id="createScheduleForm">
        <div class="form__element">
          <label for="fieldName" class="form__label">Name:</label>
          <input
            type="text"
            class="form__inp"
            id="fieldName"
            placeholder="207"
            required
          />
        </div>
        <div class="form__element">
          <div class="dropdown__wrapper dropdown-relative">
            <div class="select">
              <button
                class="btn-no-bg select__btn"
                type="button"
                name="capacity"
                value=""
                data-index="0"
                data-select="toggle"
              >
                Capacity:
                <svg
                  width="15"
                  height="16"
                  viewBox="0 0 15 16"
                  fill="none"
                  xmlns="http://www.w3.org/2000/svg"
                >
                  <path
                    d="M0.75 2.3999L7.5 9.5999L14.25 2.3999"
                    stroke="#A6A6A6"
                    stroke-width="1.5"
                    stroke-linecap="round"
                    stroke-linejoin="round"
                  />
                </svg>
              </button>
              <div class="select__dropdown">
                <ul class="select__options" data-index="-1">
                  <li
                    class="select__option"
                    data-index="0"
                    data-value="0"
                    data-select="option"
                  >
                    0
                  </li>
                  <li
                    class="select__option"
                    data-index="1"
                    data-value="1"
                    data-select="option"
                  >
                    1
                  </li>
                </ul>
              </div>
            </div>
          </div>
        </div>
        <div class="form__element">
          <div class="dropdown__wrapper dropdown-relative">
            <div class="select">
              <button
                class="btn-no-bg select__btn"
                type="button"
                name="type"
                value=""
                data-index="0"
                data-select="toggle"
              >
                Type:
                <svg
                  width="15"
                  height="16"
                  viewBox="0 0 15 16"
                  fill="none"
                  xmlns="http://www.w3.org/2000/svg"
                >
                  <path
                    d="M0.75 2.3999L7.5 9.5999L14.25 2.3999"
                    stroke="#A6A6A6"
                    stroke-width="1.5"
                    stroke-linecap="round"
                    stroke-linejoin="round"
                  />
                </svg>
              </button>
              <div class="select__dropdown">
                <ul class="select__options" data-index="0">
                  <li
                    class="select__option"
                    data-index="1"
                    data-value="default"
                    data-select="option"
                  >
                    default
                  </li>
                  <li
                    class="select__option"
                    data-index="2"
                    data-value="sport"
                    data-select="option"
                  >
                    sport
                  </li>
                  <li
                    class="select__option"
                    data-index="3"
                    data-value="computered"
                    data-select="option"
                  >
                    computered
                  </li>
                  <li
                    class="select__option"
                    data-index="4"
                    data-value="chemistry"
                    data-select="option"
                  >
                    chemistry
                  </li>
                </ul>
              </div>
            </div>
          </div>
        </div>
      </form>
    </div>
    <div class="popup__footer">
      <button id="commit" class="btn" method="post">Создать</button>
    </div>
  </div>
</template>
