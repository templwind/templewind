class Inputchip extends HTMLElement {
	constructor() {
		super();
	}

	connectedCallback(): void {
		console.log("Inputchip connected");
	}
}

customElements.define("tw-inputchip", Inputchip);