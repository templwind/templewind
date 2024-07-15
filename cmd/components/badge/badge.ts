class Badge extends HTMLElement {
	constructor() {
		super();
	}

	connectedCallback(): void {
		console.log("Badge connected");
	}
}

customElements.define("tw-badge", Badge);