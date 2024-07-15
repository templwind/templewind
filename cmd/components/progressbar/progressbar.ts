class Progressbar extends HTMLElement {
	constructor() {
		super();
	}

	connectedCallback(): void {
		console.log("Progressbar connected");
	}
}

customElements.define("tw-progressbar", Progressbar);