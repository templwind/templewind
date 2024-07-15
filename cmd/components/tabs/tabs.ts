class Tabs extends HTMLElement {
	constructor() {
		super();
	}

	connectedCallback(): void {
		console.log("Tabs connected");
	}
}

customElements.define("tw-tabs", Tabs);