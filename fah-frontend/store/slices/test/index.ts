import { createSlice, PayloadAction } from "@reduxjs/toolkit";

interface CounterState {
	count: number;
}

const state: CounterState = {
	count: 0,
}

const counterSlice = createSlice({
	name: "counter",
	initialState: state,
	reducers: {
		increment(state) {
			state.count += 1;
		},

		decrement(state) {
			state.count -= 1;
		},

		setCounter(state, action: PayloadAction<number>) {
			state.count = action.payload;
		}
	}
});

export const { increment, decrement, setCounter } = counterSlice.actions;
export default counterSlice.reducer;