<template>
  <div class="popup popup-squared">
    <div class="popup__header">
      <h2 class="h2">Редактировать выделенную область</h2>
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

    <div class="headings sticky">
      <h3 class="h3">Выбранные элементы:</h3>
      <h3 class="h3">Новые значения полей:</h3>
    </div>

    <div class="popup__main">
      <div class="pairs">
        <div class="pair">
          <div class="leftside">
            <span class="text h4">Выбранный элемент:</span>
            <div class="card">
              <div class="info-str">
                <span class="text">ID: </span
                ><span class="text-warning">100</span>
              </div>
              <div class="info-str">
                <span class="text">Capacity: </span
                ><span class="text-warning">0</span>
              </div>
              <div class="info-str">
                <span class="text">Name: </span
                ><span class="text-warning">210</span>
              </div>
              <div class="info-str">
                <span class="text">Type: </span
                ><span class="text-warning">Default</span>
              </div>
            </div>
          </div>
          <div class="rightside">
            <form action="" class="form" id="editScheduleForm">
              <div class="form__element">
                <span class="text h4">Name:</span>
                <input
                  type="text"
                  class="form__inp"
                  id="fieldName"
                  placeholder="207"
                  required
                />
              </div>
              <div class="form__element">
                <span class="text h4">Capacity:</span>
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
                      Capacity
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
                <span class="text h4">Type:</span>
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
                      Type
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
        </div>

        <div class="pair">
          <div class="leftside">
            <span class="text h4">Выбранный элемент:</span>
            <div class="card">
              <div class="info-str">
                <span class="text">ID: </span
                ><span class="text-warning">100</span>
              </div>
              <div class="info-str">
                <span class="text">Capacity: </span
                ><span class="text-warning">0</span>
              </div>
              <div class="info-str">
                <span class="text">Name: </span
                ><span class="text-warning">210</span>
              </div>
              <div class="info-str">
                <span class="text">Type: </span
                ><span class="text-warning">Default</span>
              </div>
            </div>
          </div>
          <div class="rightside">
            <form action="" class="form" id="editScheduleForm">
              <div class="form__element">
                <span class="text h4">Name:</span>
                <input
                  type="text"
                  class="form__inp"
                  id="fieldName"
                  placeholder="207"
                  required
                />
              </div>
              <div class="form__element">
                <span class="text h4">Capacity:</span>
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
                      Capacity
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
                <span class="text h4">Type:</span>
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
                      Type
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
        </div>
      </div>
    </div>

    <div class="popup__footer">
      <div class="button-group">
        <button
          class="btn btn-no-bg btn-bordered btn-bordered__red"
          type="submit"
        >
          Отменить изменения
        </button>
        <button class="btn" type="submit">Применить</button>
      </div>
    </div>
  </div>
</template>
