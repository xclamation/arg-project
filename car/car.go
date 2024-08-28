package car

import (
	"fmt"
	"strconv"
	"strings"
)

type Car struct {
	Brand string
	Speed int
}

func (c *Car) Set(value string) error {
	parts := strings.Split(value, ":")
	if len(parts) != 2 {
		return fmt.Errorf("invalid format, expected Brand:Speed")
	}

	c.Brand = parts[0]

	speed, err := strconv.Atoi(parts[1])
	if err != nil {
		return fmt.Errorf("invalid speed: %v", err)
	}

	c.Speed = speed

	return nil
}

func (c *Car) String() string {
	return fmt.Sprintf("%s:%d", c.Brand, c.Speed)
}
