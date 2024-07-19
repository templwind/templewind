import './list-group.scss';

export class TwListGroup extends HTMLElement {
	constructor() {
		super();
	}

	connectedCallback(): void {
		console.log("ListGroup connected");
	}
}

customElements.define("tw-list-group", TwListGroup);