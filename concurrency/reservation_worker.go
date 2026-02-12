package concurrency


type ReservationRequest struct {
	BookID   int
	MemberID int
	Result   chan error
}