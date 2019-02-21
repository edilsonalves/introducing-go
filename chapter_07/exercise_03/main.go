type Shape interface {
    perimeter() float64
}

func (c *Circle) perimeter() float64 {
    return 2 * math.Pi * c.r
}

func (r *Rectangle) perimeter() float64 {
    length := distance(r.x1, r.y1, r.x1, r.y2)
    width := distance(r.x1, r.y1, r.x2, r.y1)
    return 2 * (length + width)
}
