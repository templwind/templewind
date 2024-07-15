class Link extends HTMLElement {
	constructor() {
		super();
	}

	connectedCallback(): void {
		console.log("Link connected");
	}
}

customElements.define("tw-link", Link);