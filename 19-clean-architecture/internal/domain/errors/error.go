package errors

type EventError string

func (ee EventError) Error() string {
	return string(ee)
}

var (
	ErrOverlaping       = EventError("another event exists for this date")
	ErrIncorrectEndDate = EventError("end_date is incorrect")
)
