package asteroidcollision

// Collide returns the asteroids that survive. Time O(n), space O(n).
func Collide(a []int) []int {
	st := []int{}
	for _, x := range a {
		alive := true
		for alive && x < 0 && len(st) > 0 && st[len(st)-1] > 0 {
			top := st[len(st)-1]
			switch {
			case top < -x:
				st = st[:len(st)-1] // top explodes, keep checking
			case top == -x:
				st = st[:len(st)-1] // both explode
				alive = false
			default:
				alive = false // x explodes
			}
		}
		if alive {
			st = append(st, x)
		}
	}
	return st
}
