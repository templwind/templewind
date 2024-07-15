class Datepicker extends HTMLElement {
	constructor() {
		super();
	}

	connectedCallback(): void {
		console.log("Datepicker connected");
	}
}

customElements.define("tw-datepicker", Datepicker);