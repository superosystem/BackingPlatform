package com.gusrylmubarok.spring;

import javax.persistence.*;

@Entity
@Table(name="USERS")
public class User {

    @Id
    @SequenceGenerator(name="user_id_generator", sequenceName = "user_id_sequence", initialValue = 4)
    @GeneratedValue(generator = "user_id_generator")
    private Integer id;

    @Column(nullable = false)
    private String name;

    @Column(nullable = false)
    private String email;

    private boolean disable;

    public User() {
    }

    public User(Integer id, String name, String email, boolean disable) {
        this.id = id;
        this.name = name;
        this.email = email;
        this.disable = disable;
    }

    public Integer getId() {
        return id;
    }

    public void setId(Integer id) {
        this.id = id;
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public String getEmail() {
        return email;
    }

    public void setEmail(String email) {
        this.email = email;
    }

    public boolean isDisable() {
        return disable;
    }

    public void setDisable(boolean disable) {
        this.disable = disable;
    }
}
