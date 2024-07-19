import './progressbar.scss';

export class TwProgressbar extends HTMLElement {
	constructor() {
		super();
	}

	connectedCallback(): void {
		console.log("Progressbar connected");
	}
}

customElements.define("tw-progressbar", TwProgressbar);