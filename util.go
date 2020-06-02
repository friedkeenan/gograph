package main

import "image/color"

type RGBA16 struct {
    r, g, b, a uint16
}

func (c RGBA16) RGBA() (r, g, b, a uint32) {
    r = uint32(c.r)
    g = uint32(c.g)
    b = uint32(c.b)
    a = uint32(c.a)

    return
}

func MinInt(a, b int) int {
    if a <= b {
        return a
    }

    return b
}

func BlendColor(old, new color.Color) color.Color {
    old_r, old_g, old_b, _ := old.RGBA()
    new_r, new_g, new_b, new_a := new.RGBA()

    return RGBA16{
        uint16((new_a * new_r + (0xFFFF - new_a) * old_r) / 0xFFFF),
        uint16((new_a * new_g + (0xFFFF - new_a) * old_g) / 0xFFFF),
        uint16((new_a * new_b + (0xFFFF - new_a) * old_b) / 0xFFFF),
        0xFFFF,
    }
}