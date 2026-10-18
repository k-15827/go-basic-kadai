package main

import (
	"errors"
	"fmt"
)

type ValidationError struct {
	Message string  // エラーメッセージ
	Field   string  // 不正なフィールド
	Value   float64 // 不正な値
}

func (v *ValidationError) Error() string {
	return v.Message
}

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (r Rectangle) Perimeter() float64 {
	return r.Height*2 + r.Width*2
}

func (c Circle) Area() float64 {
	return c.Radius * c.Radius * 3.14
}

func (c Circle) Perimeter() float64 {
	return c.Radius * 2 * 3.14
}

func validateShape(shape Shape) error {
	switch val := shape.(type) {
	case Rectangle:
		if val.Height <= 0 {
			err := &ValidationError{
				Message: "長方形の幅は正の値としてください。",
				Field:   "Rectangle.Height",
				Value:   val.Height,
			}
			return err
		}
		if val.Width <= 0 {
			err := &ValidationError{
				Message: "長方形の横幅は正の値としてください。",
				Field:   "Rectangle.Width",
				Value:   val.Width,
			}
			return err
		}

	case Circle:
		if val.Radius < 0 {
			err := &ValidationError{
				Message: "円の半径は正の値としてください。",
				Field:   "Circle.Radius",
				Value:   val.Radius,
			}
			return err
		}
	}
	return nil
}

func main() {
	// 処理対象の図形スライスを作成
	shapes := []Shape{
		Rectangle{Width: 10.0, Height: 5.0},
		Circle{Radius: 3.0},
		Rectangle{Width: -2.0, Height: 5.0}, // 不正（幅が不正）
		Circle{Radius: -1.0},                // 不正（半径が不正）
		Rectangle{Width: 1.0, Height: 0.0},  // 不正（高さが不正）
	}

	fmt.Println("--- 図形のバリデーション ---")

	// バリデーションOKの場合：計算対象スライスcalcShapes[]にその図形を追加
	var calcShapes []Shape

	for i, s := range shapes {
		shapeErr := validateShape(s)

		if shapeErr != nil {
			var validErr *ValidationError
			if errors.As(shapeErr, &validErr) {
				fmt.Printf("#%d：%T => NG：[バリデーションエラー] %s ：%s (Value：%.2f)\n", i+1, s, validErr.Field, validErr.Message, validErr.Value)
				fmt.Printf(" [詳細] Field=%s, Value=%.2f\n", validErr.Field, validErr.Value)
			}
		} else {
			fmt.Printf("#%d：%T => バリデーションOK\n", i+1, s)
			calcShapes = append(calcShapes, s)
		}
	}

	fmt.Println("\n--- 検証済み図形の一括計算 ---")
	if len(calcShapes) != 0 {
		for _, s := range calcShapes {
			fmt.Printf("図形：%T | 面積：%.2f |周長：%.2f\n", s, s.Area(), s.Perimeter())
		}
	}
}
