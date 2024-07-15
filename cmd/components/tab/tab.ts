class Tab extends HTMLElement {
	constructor() {
		super();
	}

	connectedCallback(): void {
		console.log("Tab connected");
	}
}

customElements.define("tw-tab", Tab);