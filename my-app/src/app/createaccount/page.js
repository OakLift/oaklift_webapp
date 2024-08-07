// import {Input} from "@nextui-org/react";
import React, { useState } from 'react';
import {Button} from  '@/app/components/Button';

export default function Login() {


    return (
        <div>
            {/* <Input type='firstname' label='First Name' /> */}
            <input 
            type='firstName'
            name='First Name'
            placeholder='First Name'
            />
            <input
            type='lastName'
            name='Last Name'
            placeholder='Last Name'
            />

            <Button>Create</Button>
            
        </div>
        

    );
}