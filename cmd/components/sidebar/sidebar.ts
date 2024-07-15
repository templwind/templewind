class Sidebar extends HTMLElement {
	constructor() {
		super();
	}

	connectedCallback(): void {
		console.log("Sidebar connected");
	}
}

customElements.define("tw-sidebar", Sidebar);