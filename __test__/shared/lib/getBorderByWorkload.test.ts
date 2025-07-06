// eslint-disable-next-line @typescript-eslint/no-require-imports
require("@testing-library/jest-dom");
import { getBorderColorByWorkload } from "@/shared/lib/getBorderColorByWorkload";

describe("getBorderByWorkload", () => {
	it("returns red for overload faculty", () => {
		expect(getBorderColorByWorkload(1.01)).toBe("#FF0000");
		expect(getBorderColorByWorkload(1.15)).toBe("#FF0000");
	})

	it("returns red for workload > 0.8", () => {
		expect(getBorderColorByWorkload(0.81)).toBe("#FF0000");
		expect(getBorderColorByWorkload(1)).toBe("#FF0000");
	})

	it("returns orange for 0.65 < workload <= 0.8", () => {
		expect(getBorderColorByWorkload(0.66)).toBe('#FF9D00');
		expect(getBorderColorByWorkload(0.8)).toBe('#FF9D00');
	})

	it('returns yellow for 0.4 < workload <= 0.65', () => {
		expect(getBorderColorByWorkload(0.41)).toBe('#FFFB00');
		expect(getBorderColorByWorkload(0.65)).toBe('#FFFB00');
	});

	it('returns green for workload <= 0.4', () => {
		expect(getBorderColorByWorkload(0.4)).toBe('#88FF00');
		expect(getBorderColorByWorkload(0)).toBe('#88FF00');
	});
})