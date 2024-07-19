import './inputchip.scss';

export class TwInputchip extends HTMLElement {
	constructor() {
		super();
	}

	connectedCallback(): void {
		console.log("Inputchip connected");
	}
}

customElements.define("tw-inputchip", TwInputchip);