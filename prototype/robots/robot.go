package robots

type Robot interface {
	Execute()
	Clone() Robot
}
