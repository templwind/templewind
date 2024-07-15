class Timeline extends HTMLElement {
	constructor() {
		super();
	}

	connectedCallback(): void {
		console.log("Timeline connected");
	}
}

customElements.define("tw-timeline", Timeline);