// Copyright 2026 Carlos Munoz and the Folio Authors
// SPDX-License-Identifier: Apache-2.0

package layout

import (
	"testing"
)

func TestRenderLinearGradientBasic(t *testing.T) {
	stops := []GradientStop{
		{Color: RGB(1, 0, 0), Position: 0}, // red
		{Color: RGB(0, 0, 1), Position: 1}, // blue
	}
	img := RenderLinearGradient(100, 100, 0, stops) // 0° = to top
	if img.Bounds().Dx() != 100 || img.Bounds().Dy() != 100 {
		t.Errorf("expected 100x100, got %dx%d", img.Bounds().Dx(), img.Bounds().Dy())
	}
	// Top-center should be more blue, bottom-center more red (0° = to top).
	topR, _, topB, _ := img.At(50, 0).RGBA()
	botR, _, botB, _ := img.At(50, 99).RGBA()
	if topB <= topR {
		t.Errorf("top pixel should be more blue than red: R=%d B=%d", topR>>8, topB>>8)
	}
	if botR <= botB {
		t.Errorf("bottom pixel should be more red than blue: R=%d B=%d", botR>>8, botB>>8)
	}
}

func TestRenderLinearGradient90(t *testing.T) {
	stops := []GradientStop{
		{Color: RGB(1, 0, 0), Position: 0},
		{Color: RGB(0, 1, 0), Position: 1},
	}
	img := RenderLinearGradient(100, 50, 90, stops) // 90° = to right
	// Left should be more red, right more green.
	leftR, leftG, _, _ := img.At(0, 25).RGBA()
	rightR, rightG, _, _ := img.At(99, 25).RGBA()
	if leftR <= leftG {
		t.Errorf("left pixel should be more red: R=%d G=%d", leftR>>8, leftG>>8)
	}
	if rightG <= rightR {
		t.Errorf("right pixel should be more green: R=%d G=%d", rightR>>8, rightG>>8)
	}
}

func TestRenderLinearGradientSingleStop(t *testing.T) {
	stops := []GradientStop{
		{Color: RGB(0.5, 0.5, 0.5), Position: 0},
	}
	// Less than 2 stops → returns 1x1 fallback.
	img := RenderLinearGradient(100, 100, 0, stops)
	if img.Bounds().Dx() != 1 || img.Bounds().Dy() != 1 {
		t.Errorf("expected 1x1 fallback for single stop, got %dx%d", img.Bounds().Dx(), img.Bounds().Dy())
	}
}

func TestRenderLinearGradientDimensions(t *testing.T) {
	stops := []GradientStop{
		{Color: RGB(1, 0, 0), Position: 0},
		{Color: RGB(0, 0, 1), Position: 1},
	}
	img := RenderLinearGradient(200, 50, 45, stops)
	if img.Bounds().Dx() != 200 || img.Bounds().Dy() != 50 {
		t.Errorf("expected 200x50, got %dx%d", img.Bounds().Dx(), img.Bounds().Dy())
	}
}

func TestRenderLinearGradientZeroSize(t *testing.T) {
	stops := []GradientStop{
		{Color: RGB(1, 0, 0), Position: 0},
		{Color: RGB(0, 0, 1), Position: 1},
	}
	// Zero dimensions → 1x1 fallback.
	img := RenderLinearGradient(0, 0, 0, stops)
	if img == nil {
		t.Fatal("expected non-nil image even for zero size")
	}
}

func TestRenderLinearGradientNoStops(t *testing.T) {
	img := RenderLinearGradient(100, 100, 0, nil)
	if img == nil {
		t.Fatal("expected non-nil image for nil stops")
	}
}
