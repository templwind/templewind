import './button-group.scss';

export class TwButtonGroup extends HTMLElement {
	constructor() {
		super();
	}

	connectedCallback(): void {
		console.log("ButtonGroup connected");
	}
}

customElements.define("tw-button-group", TwButtonGroup);