package com.gusrylmubarok.demoapi.services;

import java.util.List;
import java.util.Optional;

import javax.transaction.Transactional;

import com.gusrylmubarok.demoapi.models.entity.Product;
import com.gusrylmubarok.demoapi.models.entity.Supplier;
import com.gusrylmubarok.demoapi.models.repository.ProductRepository;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

@Service
@Transactional
public class ProductService {
    
    @Autowired
    private ProductRepository productRepository;

    public Product save(Product product) {
        return productRepository.save(product);
    }

    public Product findOne(Long id) {
        Optional<Product> product = productRepository.findById(id);
        if(!product.isPresent()) {
            return null;
        }
        return product.get();
    }

    public Iterable<Product> findAll() {
        return productRepository.findAll();
    }

    public void deleteById(Long id) {
        productRepository.deleteById(id);
    }

    public List<Product> findByName(String name) {
        return productRepository.findByNameContains(name);
    }

    public void addSupplier(Supplier supplier, Long productId) {
        Product product = findOne(productId);
        if(product == null) {
            throw new RuntimeException("product with ID: " +productId+ " not found");
        }

        product.getSuppliers().add(supplier);
        save(product);
    }
}
