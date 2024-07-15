class Select extends HTMLElement {
	constructor() {
		super();
	}

	connectedCallback(): void {
		console.log("Select connected");
	}
}

customElements.define("tw-select", Select);