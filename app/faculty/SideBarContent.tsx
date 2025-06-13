import Image from "next/image";
import arrowSvg from "@/public/icons/svg/arrow.svg";
import styles from "./styles.module.scss";

const SideBarContent: React.FC = () => {
	return <div className={ styles.sideBar }>
		<div className={ styles.menu }>
			<button className={ styles.button }>
				<span>Year of study</span>
				<Image src={ arrowSvg } alt=""/>
			</button>
			<form action="1">
				<label htmlFor="1">
					<input type="checkbox" id="1"/>
					<span className={ styles.text }>BS - Year 1</span>
				</label>
				<label htmlFor="2">
					<input type="checkbox" id={ "2" }/>
					<span className={ styles.text }>BS - Year 2</span>
				</label>
				<label htmlFor="3">
					<input type="checkbox" id={ "3" }/>
					<span className={ styles.text }>BS - Year 3</span>
				</label>
				<label htmlFor="4">
					<input type="checkbox" id={ "4" }/>
					<span className={ styles.text }>MS</span>
				</label>
				<label htmlFor="5">
					<input type="checkbox" id={ "5" }/>
					<span className={ styles.text }>PhD</span>
				</label>
			</form>
			<button className={ styles.button }>
				<span>Study track</span>
				<Image src={ arrowSvg } alt=""/>
			</button>
			<form action="2">
				<label htmlFor="ISE">
					<input type="checkbox" id="ISE"/>
					<span className={ styles.text }>ISE</span>
				</label>
				<label htmlFor="DSAI">
					<input type="checkbox" id={ "DSAI" }/>
					<span className={ styles.text }>DSAI</span>
				</label>
				<label htmlFor="MFAI">
					<input type="checkbox" id={ "MFAI" }/>
					<span className={ styles.text }>MFAI</span>
				</label>
				<label htmlFor="AI360">
					<input type="checkbox" id={ "AI360" }/>
					<span className={ styles.text }>AI360</span>
				</label>
				<label htmlFor="RO">
					<input type="checkbox" id={ "RO" }/>
					<span className={ styles.text }>RO</span>
				</label>
				<label htmlFor="SE">
					<input type="checkbox" id={ "SE" }/>
					<span className={ styles.text }>SE</span>
				</label>
				<label htmlFor="RO">
					<input type="checkbox" id={ "RO" }/>
					<span className={ styles.text }>RO</span>
				</label>
				<label htmlFor="DS">
					<input type="checkbox" id={ "DS" }/>
					<span className={ styles.text }>DS</span>
				</label>
				<label htmlFor="TE">
					<input type="checkbox" id={ "TE" }/>
					<span className={ styles.text }>TE</span>
				</label>
				<label htmlFor="SNE">
					<input type="checkbox" id={ "SNE" }/>
					<span className={ styles.text }>SNE</span>
				</label>
			</form>
		</div>
	</div>
}

export default SideBarContent;