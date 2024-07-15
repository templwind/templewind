class Footer extends HTMLElement {
	constructor() {
		super();
	}

	connectedCallback(): void {
		console.log("Footer connected");
	}
}

customElements.define("tw-footer", Footer);