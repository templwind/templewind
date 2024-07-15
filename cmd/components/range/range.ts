class Range extends HTMLElement {
	constructor() {
		super();
	}

	connectedCallback(): void {
		console.log("Range connected");
	}
}

customElements.define("tw-range", Range);