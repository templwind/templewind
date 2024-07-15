class Progressradial extends HTMLElement {
	constructor() {
		super();
	}

	connectedCallback(): void {
		console.log("Progressradial connected");
	}
}

customElements.define("tw-progressradial", Progressradial);