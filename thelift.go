package kata

var queues [][]int

func TheLift(queinp [][]int, cap int) []int {
	queues = queinp

	lift := newcall(cap)

	for !nomoreppl(lift) {
		lift.move()
	}

	return lift.movements
}

func nomoreppl(lift *Lift) bool {
	for i := 0; i < len(queues); i++ {
		for j := 0; j < len(queues[i]); j++ {
			if queues[i][j] != i {
				return false
			}
		}
	}
	return len(lift.ppl) == 0 && lift.fl == 0
}

func push(person, fl int) {
	people := append(queues[fl], person)
	queues[fl] = people
}

func pop(fl, index int) int {
	people := queues[fl]
	var person int
	if len(people) == 0 {
		person = -1
	} else {
		person = people[index]
		people[index] = -1
	}
	queues[fl] = filter(people)
	return person
}

func filter(people []int) []int {
	var filteredPeople []int
	for _, val := range people {
		if val >= 0 {
			filteredPeople = append(filteredPeople, val)
		}
	}
	return filteredPeople
}

type Lift struct {
	cap       int
	ppl       []int
	movements []int
	fl        int
	travelup  bool
}

func newcall(cap int) *Lift {
	lift := &Lift{
		cap:       cap,
		ppl:       []int{},
		movements: []int{},
		fl:        -1,
		travelup:  true,
	}
	lift.swithcfl(0)
	return lift
}

func (lift *Lift) move() {
	lift.dropppl()
	lift.pickppl()
	lift.gonxtfl()
}

func (lift *Lift) dropppl() {
	var arrivedppl []int
	for _, dest := range lift.ppl {
		if lift.fl == dest {
			arrivedppl = append(arrivedppl, dest)
		}
	}
	lift.ppl = removeppl(lift.ppl, arrivedppl)
	for _, passenger := range arrivedppl {
		push(passenger, lift.fl)
	}
}

func (lift *Lift) pickppl() {
	for i := 0; i < len(queues[lift.fl]) && lift.cap > len(lift.ppl); i++ {
		dest := queues[lift.fl][i]
		if lift.isGoingSameWay(dest, lift.fl) {
			lift.ppl = append(lift.ppl, pop(lift.fl, i))
			i-- // decrease because of in use change of array
		}
	}
}

func (lift *Lift) gonxtfl() {
	if lift.travelup {
		if lift.moveUpIfUseful() {
			return
		}
		if len(lift.ppl) == 0 {
			if lift.smartUp() {
				return
			}
			lift.travelup = false
			if lift.moveDownIfUseful() {
				return
			}
			if lift.smartDown() {
				return
			}
		}
	} else {
		if lift.moveDownIfUseful() {
			return
		}
		if len(lift.ppl) == 0 {
			if lift.smartDown() {
				return
			}
			lift.travelup = true
			if lift.moveUpIfUseful() {
				return
			}
			if lift.smartUp() {
				return
			}
		}
	}
	lift.swithcfl(0)
}

func (lift *Lift) smartDown() bool {
	for i := 0; i <= lift.fl; i++ {
		if lift.moveIfAnyGoingDifferentWay(i) {
			return true
		}
	}
	return false
}

func (lift *Lift) smartUp() bool {
	for i := len(queues) - 1; i >= lift.fl; i-- {
		if lift.moveIfAnyGoingDifferentWay(i) {
			return true
		}
	}
	return false
}

func (lift *Lift) moveDownIfUseful() bool {
	for i := lift.fl - 1; i >= 0; i-- {
		if lift.moveIfPassengerdest(i) {
			return true
		}
		if lift.moveIfAnyGoingSameWay(i) {
			return true
		}
	}
	return false
}

func (lift *Lift) moveUpIfUseful() bool {
	for i := lift.fl + 1; i < len(queues); i++ {
		if lift.moveIfPassengerdest(i) {
			return true
		}
		if lift.moveIfAnyGoingSameWay(i) {
			return true
		}
	}
	return false
}

func (lift *Lift) moveIfPassengerdest(fl int) bool {
	if contains(lift.ppl, fl) {
		lift.swithcfl(fl)
		return true
	}
	return false
}

func (lift *Lift) moveIfAnyGoingSameWay(fl int) bool {
	for i := 0; i < len(queues[fl]); i++ {
		if lift.isGoingSameWay(queues[fl][i], fl) {
			lift.swithcfl(fl)
			return true
		}
	}
	return false
}

func (lift *Lift) moveIfAnyGoingDifferentWay(fl int) bool {
	for i := 0; i < len(queues[fl]); i++ {
		if queues[fl][i] != fl && !lift.isGoingSameWay(queues[fl][i], fl) {
			lift.swithcfl(fl)
			lift.travelup = !lift.travelup
			return true
		}
	}
	return false
}

func (lift *Lift) isGoingSameWay(dest, currentfl int) bool {
	if lift.travelup {
		return dest > currentfl
	}
	return dest < currentfl
}

func (lift *Lift) swithcfl(fl int) {
	if lift.fl == fl {
		return
	}
	if lift.travelup && fl < lift.fl {
		lift.travelup = false
	} else if !lift.travelup && fl > lift.fl {
		lift.travelup = true
	}
	lift.fl = fl
	lift.movements = append(lift.movements, fl)
	if fl == len(queues)-1 {
		lift.travelup = false
	}
	if fl == 0 {
		lift.travelup = true
	}
}

func removeppl(slice []int, items []int) []int {
	var result []int
	for _, v := range slice {
		if !contains(items, v) {
			result = append(result, v)
		}
	}
	return result
}

func contains(slice []int, value int) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}
