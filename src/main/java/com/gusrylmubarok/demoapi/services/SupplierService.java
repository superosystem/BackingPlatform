package com.gusrylmubarok.demoapi.services;

import java.util.Optional;

import javax.transaction.TransactionScoped;

import com.gusrylmubarok.demoapi.models.entity.Supplier;
import com.gusrylmubarok.demoapi.models.repository.SupplierRepository;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

@Service
@TransactionScoped
public class SupplierService {

  @Autowired
  private SupplierRepository supplierRepository;

  public Supplier save(Supplier supplier) {
    return supplierRepository.save(supplier);
  }

  public Supplier findOne(Long id) {
    Optional<Supplier> supplier = supplierRepository.findById(id);
    if(!supplier.isPresent()) {
      return null;
    }

    return supplier.get();
  }

  public Iterable<Supplier> findAll() {
    return supplierRepository.findAll();
  }
  
  public void removeOne(Long id) {
    supplierRepository.deleteById(id);
  }
  
}
