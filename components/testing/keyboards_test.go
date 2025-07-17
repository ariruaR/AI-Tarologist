package testing

import (
	"bot/components/keyboards"
	"testing"
)

func TestCreateBuyKeyboard(t *testing.T) {
	keyboard := keyboards.CreateBuyKeyboard()

	if keyboard == nil {
		t.Fatal("CreateBuyKeyboard() returned nil")
	}

	// Проверяем, что клавиатура создается корректно
	if !keyboard.ResizeKeyboard {
		t.Error("expected ResizeKeyboard to be true")
	}
}

func TestKeyboardButtons(t *testing.T) {
	// Проверяем, что кнопки определены (они не могут быть nil, так как это структуры)
	// Просто проверяем, что код компилируется и кнопки доступны
	_ = keyboards.BtnStarCard
	_ = keyboards.BtnNotalCard
	_ = keyboards.BtnAnotherBuy

	// Тест проходит, если кнопки определены корректно
}

func TestKeyboardConsistency(t *testing.T) {
	keyboard1 := keyboards.CreateBuyKeyboard()
	keyboard2 := keyboards.CreateBuyKeyboard()

	// Проверяем, что клавиатуры создаются консистентно
	if keyboard1.ResizeKeyboard != keyboard2.ResizeKeyboard {
		t.Error("keyboards are not consistent")
	}
}
