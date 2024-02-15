package w4

// The callaback function to call on the game start.
var Start func()

// The callback function to call for each game frame refresh.
var Update func()

// //go:export start
// func start() {
// 	if Start != nil {
// 		Start()
// 	}
// }

// //go:export update
// func update() {
// 	if Update != nil {
// 		Update()
// 	}
// }
