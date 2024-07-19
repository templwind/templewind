import './avatar.scss';

export class TwAvatar extends HTMLElement {
	constructor() {
		super();
	}

	connectedCallback(): void {
		console.log("Avatar connected");
	}
}

customElements.define("tw-avatar", TwAvatar);