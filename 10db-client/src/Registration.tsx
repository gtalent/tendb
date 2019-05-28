
import React from 'react';
import Button from '@material-ui/core/Button';
import Divider from '@material-ui/core/Divider';
import TextField from '@material-ui/core/TextField';

class Registration extends React.Component {

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
			height: 187,
			width: 325,
			display: 'inline-block',
			textAlign: 'center',
		};
		return (
			<div style={style}>
				<TextField label='First Name' onChange={this.makeHandler("firstName")}/>
				<TextField label='Last Name' onChange={this.makeHandler("lastName")}/>
				<TextField label='Birthday' onChange={this.makeHandler("birthday")} type='date' InputLabelProps={{shrink: true}}/>
				<TextField label='Email' onChange={this.handleEmailChange}/>
				<Divider/>
				<TextField label='Password'
							  type='password'
							  onChange={this.handlePasswordChange}
				/>
				<Divider/>
				<TextField label='Confirm Password'
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
