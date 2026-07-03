// ... existing code ...

func (cs *clientStream) newAttempt(ctx context.Context) *csAttempt {
	// Ensure the context used for the attempt is the original context
	// which contains the deadline and metadata.
	return &csAttempt{
		ctx: ctx,
		// ... existing fields ...
	}
}

// ... existing code ...