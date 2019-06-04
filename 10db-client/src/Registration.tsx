
import React from 'react';
import Button from '@material-ui/core/Button';
import Divider from '@material-ui/core/Divider';
import TextField from '@material-ui/core/TextField';

interface RegistrationProps {
}

interface RegistrationState {
	firstName: string;
	lastName: string;
	email: string;
	password: string;
	passwordConfirm: string;
}

class Registration extends React.Component<RegistrationProps, RegistrationState> {

	constructor(props: RegistrationProps) {
		super(props);
		this.state = {
			firstName: '',
			lastName: '',
			email: '',
			password: '',
			passwordConfirm: '',
		};
	}

	makeHandler = (param: string) => {
		return (event: any) => {
			let obj:any = {};
			obj[param] = event.target.value;
			this.setState(obj);
		};
	};

	handleEmailChange = (event: any) => {
		this.setState({
			email: event.target.value,
		});
	};

	handlePasswordChange = (event: any) => {
		this.setState({
			password: event.target.value,
		});
	};

	handlePasswordConfirmationChange = (event: any) => {
		this.setState({
			passwordConfirm: event.target.value,
		});
	};

	registerClick = (event: any) => {
	};

	render() {
		const style = {
			margin: 'auto',
			top: 0,
			bottom: 0,
			left: 0,
			right: 0,
			borderRadius: '3px',
			position: 'absolute' as 'absolute',
			height: 487,
			width: 325,
			display: 'flex',
			flexWrap: 'wrap' as 'wrap',
			justifyContent: 'center',
			alignItems: 'center',
		};
		return (
			<div style={style}>
				<TextField label='First Name' style={{width: style.width}} value={this.state.firstName || ''} onChange={this.makeHandler("firstName")}/>
				<TextField label='Last Name' style={{width: style.width}} onChange={this.makeHandler("lastName")}/>
				<TextField label='Birthday' style={{width: style.width}} onChange={this.makeHandler("birthday")} type='date' InputLabelProps={{shrink: true}}/>
				<TextField label='Email' style={{width: style.width}} onChange={this.handleEmailChange}/>
				<Divider/>
				<TextField label='Password'
				           style={{width: style.width}}
							  type='password'
							  onChange={this.handlePasswordChange}
				/>
				<Divider/>
				<TextField label='Confirm Password'
				           style={{width: style.width}}
							  type='confirmPassword'
							  onChange={this.handlePasswordConfirmationChange}
				/>
				<Divider/>
				<Button style={{width: style.width * 0.75}}
						  color='primary'
						  variant="contained"
						  onClick={this.registerClick}
				>
					 Register
				</Button>
			</div>
		);
	}

};

export default Registration;
