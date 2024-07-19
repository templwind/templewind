import './list.scss';

export class TwList extends HTMLElement {
	constructor() {
		super();
	}

	connectedCallback(): void {
		console.log("List connected");
	}
}

customElements.define("tw-list", TwList);