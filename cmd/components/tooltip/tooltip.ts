class Tooltip extends HTMLElement {
	constructor() {
		super();
	}

	connectedCallback(): void {
		console.log("Tooltip connected");
	}
}

customElements.define("tw-tooltip", Tooltip);