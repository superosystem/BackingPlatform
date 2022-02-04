package com.gusrylmubarok.demoapi.controllers;

import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.GetMapping;

@RestController
@RequestMapping("/api/welcome")
public class WelcomeController {
    
    @GetMapping
    public String welcome() {
        return "Welcome to Spring Boot - Rest API";
    }

}
