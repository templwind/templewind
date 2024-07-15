class Filebutton extends HTMLElement {
	constructor() {
		super();
	}

	connectedCallback(): void {
		console.log("Filebutton connected");
	}
}

customElements.define("tw-filebutton", Filebutton);