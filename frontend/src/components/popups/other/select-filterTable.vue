<template>
    <div class="select__dropdown dropup dropshadow">
        <ul class="select__options">
            <li class="select__option">
                <input type="radio" class="radio" id="showAll" name="filter" value="Показывать все" checked>
                <label for="showAll" class="radio-label">Показывать все</label>
            </li>
            <li class="select__option">
                <input type="radio" class="radio" id="showChecked" name="filter"
                    value="Показывать только выбранные поля">
                <label for="showChecked" class="radio-label">Показывать только выбранные поля</label>
            </li>
            <li class="select__option">
                <input type="radio" class="radio" id="showUnchecked" name="filter" value="Показывать не выбранные поля">
                <label for="showUnchecked" class="radio-label">Показывать не выбранные поля</label>
            </li>
        </ul>
    </div>
</template>
