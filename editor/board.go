package editor

import "github.com/flywave/topotypes/utils"

const (
	BOARD_LAYOUT_NONE = iota
	BOARD_LAYOUT_FILL
	BOARD_LAYOUT_FIT
	BOARD_LAYOUT_STRETCH
	BOARD_LAYOUT_TILE
	BOARD_LAYOUT_CENTER
)

func BoardTypeToString(tp int) string {
	switch tp {
	case BOARD_LAYOUT_FILL:
		return "fill"
	case BOARD_LAYOUT_FIT:
		return "fit"
	case BOARD_LAYOUT_STRETCH:
		return "stretch"
	case BOARD_LAYOUT_TILE:
		return "tile"
	case BOARD_LAYOUT_CENTER:
		return "center"
	default:
		return ""
	}
}

func StringToBoardType(tp string) int {
	if utils.StrEquals(tp, "fill") {
		return BOARD_LAYOUT_FILL
	} else if utils.StrEquals(tp, "fit") {
		return BOARD_LAYOUT_FIT
	} else if utils.StrEquals(tp, "stretch") {
		return BOARD_LAYOUT_STRETCH
	} else if utils.StrEquals(tp, "tile") {
		return BOARD_LAYOUT_TILE
	} else if utils.StrEquals(tp, "center") {
		return BOARD_LAYOUT_CENTER
	}
	return BOARD_LAYOUT_NONE
}

type Board struct {
	Layout string
	Name   string     `json:"name"`
	P1     [3]float64 `json:"p1"`
	P2     [3]float64 `json:"p2"`
	Dir    [3]float64 `json:"dir"`
}
