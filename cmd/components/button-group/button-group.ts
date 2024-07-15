class ButtonGroup extends HTMLElement {
	constructor() {
		super();
	}

	connectedCallback(): void {
		console.log("ButtonGroup connected");
	}
}

customElements.define("tw-button-group", ButtonGroup);