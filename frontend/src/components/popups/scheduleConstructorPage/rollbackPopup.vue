<template>
  <div class="popup popup-horisontal" id="rollBackPopup">
    <div class="popup__header">
      <p class="h2 text text-warning">
        Вы уверены, что хотите отменить все изменения?
      </p>
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
      <p class="text h3">
        Нажав на кнопку, вы вернете все выделенные элементы всех объектов в
        изначальное состояние.
      </p>
    </div>

    <div class="popup__footer">
      <button
        id="rollback"
        class="btn btn-bordered button btn btn-bordered btn-bordered__red"
      >
        Отменить
        <svg
          width="22"
          height="22"
          viewBox="0 0 22 22"
          fill="none"
          xmlns="http://www.w3.org/2000/svg"
        >
          <line
            x1="6.01037"
            y1="5.65674"
            x2="15.5563"
            y2="15.2027"
            stroke="#8C2B2B"
            stroke-width="1.5"
            stroke-linecap="round"
          />
          <line
            x1="5.65674"
            y1="15.2025"
            x2="15.2027"
            y2="5.65658"
            stroke="#8C2B2B"
            stroke-width="1.5"
            stroke-linecap="round"
          />
        </svg>
      </button>
    </div>
  </div>
</template>
