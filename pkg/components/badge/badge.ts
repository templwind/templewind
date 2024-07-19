import './badge.scss';

export class TwBadge extends HTMLElement {
	constructor() {
		super();
	}

	connectedCallback(): void {
		console.log("Badge connected");
	}
}

customElements.define("tw-badge", TwBadge);