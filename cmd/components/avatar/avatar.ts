class Avatar extends HTMLElement {
	constructor() {
		super();
	}

	connectedCallback(): void {
		console.log("Avatar connected");
	}
}

customElements.define("tw-avatar", Avatar);