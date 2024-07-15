class Progress extends HTMLElement {
	constructor() {
		super();
	}

	connectedCallback(): void {
		console.log("Progress connected");
	}
}

customElements.define("tw-progress", Progress);