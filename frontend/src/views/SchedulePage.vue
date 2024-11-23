<template>
  <div>
    <header class="header container header__container">
      <h1 class="h1">Расписание</h1>
      <button class="open_menu-btn btn-no-bg">
        <svg
          width="60"
          height="60"
          viewBox="0 0 60 60"
          fill="none"
          xmlns="http://www.w3.org/2000/svg"
        >
          <g clip-path="url(#clip0_262_259)">
            <path
              d="M24.4167 45.1668C24.4167 45.6531 24.6098 46.1194 24.9536 46.4632C25.2974 46.807 25.7638 47.0002 26.25 47.0002H50V43.3335H26.25C25.7638 43.3335 25.2974 43.5266 24.9536 43.8705C24.6098 44.2143 24.4167 44.6806 24.4167 45.1668Z"
              fill="black"
            />
            <path
              d="M11.5 35.1668C11.5 35.6531 11.6932 36.1194 12.037 36.4632C12.3808 36.807 12.8471 37.0002 13.3333 37.0002H50V33.3335H13.3333C13.0926 33.3335 12.8542 33.3809 12.6317 33.4731C12.4093 33.5652 12.2072 33.7002 12.037 33.8705C11.8667 34.0407 11.7317 34.2428 11.6396 34.4652C11.5474 34.6877 11.5 34.9261 11.5 35.1668Z"
              fill="black"
            />
            <path
              d="M22.3333 25.1668C22.3333 25.6531 22.5265 26.1194 22.8703 26.4632C23.2141 26.807 23.6804 27.0002 24.1667 27.0002H50V23.3335H24.1667C23.6804 23.3335 23.2141 23.5267 22.8703 23.8705C22.5265 24.2143 22.3333 24.6806 22.3333 25.1668Z"
              fill="black"
            />
            <path
              d="M11.25 13.3335C11.0092 13.3335 10.7708 13.3809 10.5484 13.4731C10.326 13.5652 10.1239 13.7002 9.95363 13.8705C9.78339 14.0407 9.64834 14.2428 9.55621 14.4652C9.46408 14.6877 9.41666 14.9261 9.41666 15.1668C9.41666 15.4076 9.46408 15.646 9.55621 15.8684C9.64834 16.0908 9.78339 16.293 9.95363 16.4632C10.1239 16.6334 10.326 16.7685 10.5484 16.8606C10.7708 16.9527 11.0092 17.0002 11.25 17.0002H50V13.3335H11.25Z"
              fill="black"
            />
          </g>
        </svg>
      </button>
    </header>
    <div class="main main__container container">
      <div class="main__left">
        <button class="btn-no-bg">
          <svg
            width="16"
            height="15"
            viewBox="0 0 16 15"
            fill="none"
            xmlns="http://www.w3.org/2000/svg"
          >
            <path
              d="M13.6001 0.75L6.4001 7.5L13.6001 14.25"
              stroke="#A6A6A6"
              stroke-width="1.5"
              stroke-linecap="round"
              stroke-linejoin="round"
            />
          </svg>
          Назад к выбору расписания
        </button>

        <div class="main__info">
          <div class="info__schedule">
            <div class="info__block">
              <h3 class="h3">Тип объекта:</h3>
              <span class="info__text text">Группы</span>
            </div>
          </div>
          <div class="info__schedule">
            <div class="info__block">
              <h3 class="h3">Неделя:</h3>
              <span class="info__text text">Чётная</span>
            </div>
          </div>
          <div class="info__schedule">
            <!-- info__block появляется в случае, если конкретный объект не выбран  -->
            <div class="info__block">
              <h3 class="h3">Объект:</h3>
              <span class="info__text text">Все объекты</span>
            </div>
          </div>
        </div>
      </div>

      <div class="main__right">
        <div class="tables">
          <div class="table__wrapper">
            <div class="info__block">
              <h3 class="h3">Объект:</h3>
              <span class="text">203-ИС-23</span>
            </div>
            <table class="table">
              <thead>
                <tr>
                  <th></th>
                  <th>ПН</th>
                  <th>ВТ</th>
                  <th>СР</th>
                  <th>ЧТ</th>
                  <th>ПТ</th>
                  <th>СБ</th>
                </tr>
              </thead>
              <tbody>
                <tr>
                  <th>1</th>
                  <td>
                    <table class="inner__table">
                      <tbody>
                        <tr>
                          <th>Кабинет:</th>
                          <td>207</td>
                        </tr>
                        <tr>
                          <th>Преподаватель:</th>
                          <td>Иванов И. И.</td>
                        </tr>
                        <tr>
                          <th>Предмет:</th>
                          <td>Физика</td>
                        </tr>
                      </tbody>
                    </table>
                  </td>
                  <td></td>
                  <td></td>
                  <td></td>
                  <td>Lorem, ipsum.</td>
                  <td></td>
                </tr>
                <tr>
                  <th>2</th>
                  <td></td>
                  <td></td>
                  <td>
                    <table class="inner__table">
                      <tbody>
                        <tr>
                          <th>Кабинет:</th>
                          <td>603</td>
                        </tr>
                        <tr>
                          <th>Преподаватель:</th>
                          <td>Иванов И. И.</td>
                        </tr>
                        <tr>
                          <th>Предмет:</th>
                          <td>ИТ</td>
                        </tr>
                      </tbody>
                    </table>
                  </td>
                  <td></td>
                  <td></td>
                  <td></td>
                </tr>
                <tr>
                  <th>3</th>
                  <td></td>
                  <td>
                    <table class="inner__table">
                      <tbody>
                        <tr>
                          <th>Кабинет:</th>
                          <td>207</td>
                        </tr>
                        <tr>
                          <th>Преподаватель:</th>
                          <td>Иванов И. И.</td>
                        </tr>
                        <tr>
                          <th>Предмет:</th>
                          <td>Психология общения</td>
                        </tr>
                      </tbody>
                    </table>
                  </td>
                  <td>Lorem ipsum dolor sit.</td>
                  <td>Lorem ipsum dolor sit.</td>
                  <td>Lorem ipsum dolor sit.</td>
                  <td></td>
                </tr>
                <tr>
                  <th>4</th>
                  <td></td>
                  <td>Lorem ipsum dolor sit amet.</td>
                  <td>Lorem ipsum dolor sit amet.</td>
                  <td></td>
                  <td></td>
                  <td>Lorem ipsum dolor sit amet.</td>
                </tr>
                <tr>
                  <th>5</th>
                  <td>Lorem ipsum dolor sit amet.</td>
                  <td></td>
                  <td>Lorem ipsum dolor sit amet.</td>
                  <td>Lorem ipsum dolor sit amet.</td>
                  <td></td>
                  <td></td>
                </tr>
                <tr>
                  <th>6</th>
                  <td></td>
                  <td></td>
                  <td></td>
                  <td></td>
                  <td>Lorem ipsum dolor sit amet.</td>
                  <td>Lorem ipsum dolor sit amet.</td>
                </tr>
                <tr>
                  <th>7</th>
                  <td>Lorem ipsum dolor sit.</td>
                  <td></td>
                  <td>Lorem ipsum dolor sit.</td>
                  <td></td>
                  <td></td>
                  <td></td>
                </tr>
              </tbody>
            </table>
          </div>
          <div class="table__wrapper">
            <div class="info__block">
              <h3 class="h3">Объект:</h3>
              <span class="text">204-ИС-23</span>
            </div>
            <table class="table">
              <thead>
                <tr>
                  <th></th>
                  <th>ПН</th>
                  <th>ВТ</th>
                  <th>СР</th>
                  <th>ЧТ</th>
                  <th>ПТ</th>
                  <th>СБ</th>
                </tr>
              </thead>
              <tbody>
                <tr>
                  <th>1</th>
                  <td>Lorem, ipsum dolor.</td>
                  <td></td>
                  <td></td>
                  <td></td>
                  <td>Lorem, ipsum.</td>
                  <td></td>
                </tr>
                <tr>
                  <th>2</th>
                  <td></td>
                  <td></td>
                  <td>Lorem, ipsum dolor.</td>
                  <td></td>
                  <td></td>
                  <td></td>
                </tr>
                <tr>
                  <th>3</th>
                  <td></td>
                  <td>Lorem ipsum dolor sit.</td>
                  <td>Lorem ipsum dolor sit.</td>
                  <td>Lorem ipsum dolor sit.</td>
                  <td>Lorem ipsum dolor sit.</td>
                  <td></td>
                </tr>
                <tr>
                  <th>4</th>
                  <td></td>
                  <td>Lorem ipsum dolor sit amet.</td>
                  <td>Lorem ipsum dolor sit amet.</td>
                  <td></td>
                  <td></td>
                  <td>Lorem ipsum dolor sit amet.</td>
                </tr>
                <tr>
                  <th>5</th>
                  <td>Lorem ipsum dolor sit amet.</td>
                  <td></td>
                  <td>Lorem ipsum dolor sit amet.</td>
                  <td>Lorem ipsum dolor sit amet.</td>
                  <td></td>
                  <td></td>
                </tr>
                <tr>
                  <th>6</th>
                  <td></td>
                  <td></td>
                  <td></td>
                  <td></td>
                  <td>Lorem ipsum dolor sit amet.</td>
                  <td>Lorem ipsum dolor sit amet.</td>
                </tr>
                <tr>
                  <th>7</th>
                  <td>Lorem ipsum dolor sit.</td>
                  <td></td>
                  <td>Lorem ipsum dolor sit.</td>
                  <td></td>
                  <td></td>
                  <td></td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>

        <button
          class="btn btn-bordered btn-bordered__blue btn-no-bg"
          id="showTableBtn"
        >
          Развернуть
          <svg
            width="52"
            height="52"
            viewBox="0 0 52 52"
            fill="none"
            xmlns="http://www.w3.org/2000/svg"
          >
            <g clip-path="url(#clip0_399_10)">
              <path
                d="M21.125 27.625H11.375C10.513 27.625 9.6864 27.9674 9.0769 28.5769C8.46741 29.1864 8.125 30.013 8.125 30.875V40.625C8.125 41.487 8.46741 42.3136 9.0769 42.9231C9.6864 43.5326 10.513 43.875 11.375 43.875H21.125C21.987 43.875 22.8136 43.5326 23.4231 42.9231C24.0326 42.3136 24.375 41.487 24.375 40.625V30.875C24.375 30.013 24.0326 29.1864 23.4231 28.5769C22.8136 27.9674 21.987 27.625 21.125 27.625ZM11.375 40.625V30.875H21.125V40.625H11.375Z"
                fill="black"
              />
              <path
                d="M30.875 34.125V37.375H40.625C41.487 37.375 42.3136 37.0326 42.9231 36.4231C43.5326 35.8136 43.875 34.987 43.875 34.125V11.375C43.875 10.513 43.5326 9.6864 42.9231 9.0769C42.3136 8.46741 41.487 8.125 40.625 8.125H17.875C17.013 8.125 16.1864 8.46741 15.5769 9.0769C14.9674 9.6864 14.625 10.513 14.625 11.375V21.125H17.875V11.375H40.625V34.125"
                fill="black"
              />
            </g>
          </svg>
        </button>
      </div>
    </div>
  </div>
</template>


<style>
@import url("@/assets/css/main.css");
@import url("@/assets/css/schedule.css");
</style>
