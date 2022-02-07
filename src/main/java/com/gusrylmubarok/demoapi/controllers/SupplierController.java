package com.gusrylmubarok.demoapi.controllers;

import javax.validation.Valid;

import com.gusrylmubarok.demoapi.dto.ResponseData;
import com.gusrylmubarok.demoapi.dto.SupplierDTO;
import com.gusrylmubarok.demoapi.models.entity.Supplier;
import com.gusrylmubarok.demoapi.services.SupplierService;

import org.modelmapper.ModelMapper;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.validation.Errors;
import org.springframework.validation.ObjectError;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.PutMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping("/api/suppliers")
public class SupplierController {
  
  @Autowired
  private SupplierService supplierService;

  @Autowired
  private ModelMapper modelMapper;

  @PostMapping
  public ResponseEntity<ResponseData<Supplier>> create(@Valid @RequestBody SupplierDTO supplierDTO, Errors errors) {
    ResponseData<Supplier> responseData = new ResponseData<>();
 
    if(errors.hasErrors()) {
      for(ObjectError error : errors.getAllErrors()) {
        responseData.getMessages().add(error.getDefaultMessage());
      }
      responseData.setStatus(false);
      responseData.setPayload(null);

      return ResponseEntity.status(HttpStatus.BAD_REQUEST).body(responseData);
    }

    Supplier supplier = modelMapper.map(supplierDTO, Supplier.class);
    
    responseData.setStatus(true);
    responseData.setPayload(supplierService.save(supplier));

    return ResponseEntity.ok(responseData);
  }

  @GetMapping
  public Iterable<Supplier> findAll() {
    return supplierService.findAll();
  }

  @GetMapping("/{id}")
  public Supplier findOne(@PathVariable("id") Long id) {
    return supplierService.findOne(id);
  }

  @PutMapping
  public ResponseEntity<ResponseData<Supplier>> update(@Valid @RequestBody SupplierDTO supplierDTO, Errors errors) {
    ResponseData<Supplier> responseData = new ResponseData<>();
 
    if(errors.hasErrors()) {
      for(ObjectError error : errors.getAllErrors()) {
        responseData.getMessages().add(error.getDefaultMessage());
      }
      responseData.setStatus(false);
      responseData.setPayload(null);

      return ResponseEntity.status(HttpStatus.BAD_REQUEST).body(responseData);
    }

    Supplier supplier = modelMapper.map(supplierDTO, Supplier.class);
    
    responseData.setStatus(true);
    responseData.setPayload(supplierService.save(supplier));

    return ResponseEntity.ok(responseData);
  } 

}
