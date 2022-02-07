package com.gusrylmubarok.demoapi.dto;

import javax.validation.constraints.NotEmpty;

public class CategoryDTO {
  
  @NotEmpty(message = "Name is required")
  private String name;

  public String getName() {
    return name;
  }

  public void setName(String name) {
    this.name = name;
  }

  
}
